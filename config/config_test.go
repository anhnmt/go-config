package config

import (
	"testing"
)

type testConfig struct {
	Name string `mapstructure:"name" default:"test"`
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
		},
		{
			name: "nil cfg",
			args: args{
				dir: "../testdata",
				env: "dev",
				cfg: nil,
			},
			wantErr: true,
		},
		{
			name: "nil pointer cfg",
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
				t.Fatalf("InitConfig() got = %v, want %v", cfg.Name, tt.wantVal)
			}
		})
	}
}
