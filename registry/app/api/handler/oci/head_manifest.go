//  Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oci

import (
	"net/http"

	"github.com/harness/gitness/registry/app/dist_temp/errcode"
	"github.com/harness/gitness/registry/app/pkg/commons"
	"github.com/harness/gitness/registry/app/pkg/docker"
	"github.com/harness/gitness/registry/app/storage"
)

// HeadManifest fetches the image manifest from the storage backend, if it exists.
func (h *Handler) HeadManifest(w http.ResponseWriter, r *http.Request) {
	info, err := h.GetRegistryInfo(r, true)
	if err != nil {
		handleErrors(r.Context(), errcode.Errors{err}, w)
		return
	}

	result := h.Controller.HeadManifest(
		r.Context(),
		info,
		r.Header[storage.HeaderAccept],
		r.Header[storage.HeaderIfNoneMatch],
	)
	if commons.IsEmpty(result.GetErrors()) {
		response, ok := result.(*docker.GetManifestResponse)
		if !ok {
			handleErrors(r.Context(), errcode.Errors{errcode.ErrCodeManifestUnknown}, w)
			return
		}
		response.ResponseHeaders.WriteToResponse(w)
	}
	handleErrors(r.Context(), result.GetErrors(), w)
}
