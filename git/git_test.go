package git

import (
	"testing"
)

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
					Enable: false,
					Url:    "https://github.com/sunny0826/config",
					Out:    "configtest",
					Token:  "xxx-",
				},
				Check: Check{
					Enable: true,
					Commit: "412c42fed6a68baf5fe2af06add37542282b8221",
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
