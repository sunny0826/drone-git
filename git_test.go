package main

import "testing"

func TestPlugin_Exec(t *testing.T) {
	type fields struct {
		Config Config
		Check  Check
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				Config: Config{
					Enable: true,
					Url:    "https://git.keking.cn/tangshd/kk-deploy-configure.git",
					Out:    "configtest",
				},
				Check: Check{
					Enable: true,
					Commit: "a838345fbbe940cfcf2678dd465f942f5408ddc3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plugin{
				Config: tt.fields.Config,
				Check:  tt.fields.Check,
			}
			if err := p.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
