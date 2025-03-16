package config

import (
	"testing"
)

type nestedConfig struct {
	SubName string `mapstructure:"sub_name" default:"nested"`
}

type testConfig struct {
	Name   string       `mapstructure:"name" default:"test"`
	PtrVal *string      `mapstructure:"ptr_val" default:"ptr_default"`
	Nested nestedConfig `mapstructure:"nested"`
}

func TestInitConfig(t *testing.T) {
	type args struct {
		dir string
		env string
		cfg any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantVal string
		wantPtr *string
		wantSub string
	}{
		{
			name: "valid config with default",
			args: args{
				dir: "../testdata",
				env: "nothing",
				cfg: &testConfig{},
			},
			wantErr: false,
			wantVal: "test",
			wantPtr: strPtr("ptr_default"),
			wantSub: "nested",
		},
		{
			name: "valid config with env override",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: &testConfig{},
			},
			wantErr: false,
			wantVal: "dev",
			wantPtr: strPtr("ptr_default"),
			wantSub: "nested",
		},
		{
			name: "config with non-existing env",
			args: args{
				dir: "../testdata",
				env: "non_existing_env",
				cfg: &testConfig{},
			},
			wantErr: false,
			wantVal: "dev",
			wantPtr: strPtr("ptr_default"),
			wantSub: "nested",
		},
		{
			name: "config with pointer field and nil",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: &testConfig{PtrVal: nil},
			},
			wantErr: false,
			wantVal: "dev",
			wantPtr: strPtr("ptr_default"), // Default value should be set
			wantSub: "nested",
		},
		{
			name: "config with nested struct",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: &testConfig{},
			},
			wantErr: false,
			wantVal: "dev",
			wantPtr: strPtr("ptr_default"),
			wantSub: "nested",
		},
		{
			name: "cfg is nil",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: nil,
			},
			wantErr: true,
		},
		{
			name: "cfg is a nil pointer",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: (*testConfig)(nil),
			},
			wantErr: true,
		},
		{
			name: "empty dir",
			args: args{
				dir: "",
				env: "err",
				cfg: nil,
			},
			wantErr: true,
		},
		{
			name: "cfg is not a struct",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: 1,
			},
			wantErr: true,
		},
		{
			name: "cfg is a struct (not pointer)",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: testConfig{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitConfig(tt.args.dir, tt.args.env, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.cfg == nil || err != nil {
				return
			}

			cfg, ok := tt.args.cfg.(*testConfig)
			if !ok {
				t.Fatalf("InitConfig() cfg is not of type *testConfig, got %T", tt.args.cfg)
			}

			if cfg.Name != tt.wantVal {
				t.Fatalf("InitConfig() got Name = %v, want %v", cfg.Name, tt.wantVal)
			}

			if cfg.PtrVal == nil || *cfg.PtrVal != *tt.wantPtr {
				t.Fatalf("InitConfig() got PtrVal = %v, want %v", cfg.PtrVal, tt.wantPtr)
			}

			if cfg.Nested.SubName != tt.wantSub {
				t.Fatalf("InitConfig() got Nested.SubName = %v, want %v", cfg.Nested.SubName, tt.wantSub)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
