package zookeeper

// func Test_zookeeperMNTR(t *testing.T) {
// 	type args struct {
// 		container types.Container
// 		tar       *tar.Tarball
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			zookeeperMNTR(tt.args.container, tt.args.tar)
// 		})
// 	}
// }

// func Test_externalIP(t *testing.T) {
// 	// This is broken. The function gets the first non-loopback interface and
// 	//  the test is looking for the default interface routing to the internet

// 	conn, _ := net.DialTimeout("udp", "10.10.10.10:80", 1*time.Millisecond)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	defer conn.Close()
// 	localAddr := conn.LocalAddr().(*net.UDPAddr)
// 	// localIP := fmt.Sprintf("%s", localAddr.IP)

// 	tests := []struct {
// 		name    string
// 		want    string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "First non loopback IP interface address",
// 			want:    localAddr.IP.String(),
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			got, err := externalIP()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("externalIP() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("externalIP() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
