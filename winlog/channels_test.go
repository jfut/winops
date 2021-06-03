// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build windows

package winlog

import (
	"testing"

	"bitbucket.org/creachadair/stringset"
)

func TestAvailableChannels(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		{
			// Exists on all Windows machines.
			name:    "must exist",
			want:    []string{"Application", "Security", "System"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AvailableChannels()
			if (err != nil) != tt.wantErr {
				t.Errorf("AvailableChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !stringset.New(got...).Contains(tt.want...) {
				t.Errorf("AvailableChannels() = %v, must contain %v", got, tt.want)
			}
		})
	}
}
