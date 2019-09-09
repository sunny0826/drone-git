package git

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
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
					//Url:    "https://github.com/sunny0826/config",
					Out: "configtest",
					//Token:  "xxx-",
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

func Test_commandDiffCommit(t *testing.T) {
	mergeCmd := commandDiffCommit()
	mergeOut, _ := mergeCmd.Output()
	result := strings.Fields(string(mergeOut))
	var list []string
	for i, n := range result {
		if n == "M" {
			list = append(list, result[i+1])
		}
	}
	fmt.Println(list)
}

func Test_commandClone(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
		{
			name:"test",
			args:args{
				config:Config{
					Url:"https://git.keking.ops/tangshd/kk-deploy-configure.git",
					Token:"test123",
					Out:"config",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandClone(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commandClone() = %v, want %v", got, tt.want)
			}
		})
	}
}
