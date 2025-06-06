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

package npm

import (
	"net/http"

	"github.com/harness/gitness/registry/app/dist_temp/errcode"
	"github.com/harness/gitness/registry/app/pkg/commons"
	"github.com/harness/gitness/registry/app/pkg/types/npm"
	"github.com/harness/gitness/registry/request"
)

func (h *handler) DeleteVersion(w http.ResponseWriter, r *http.Request) {
	contextInfo := request.ArtifactInfoFrom(r.Context())
	info, ok := contextInfo.(npm.ArtifactInfo)
	if !ok {
		h.HandleErrors2(r.Context(), errcode.ErrCodeInvalidRequest.WithMessage("failed to fetch info from context"), w)
		return
	}

	response := h.controller.DeleteVersion(r.Context(), info)
	if !commons.IsEmpty(response.GetError()) {
		http.Error(w, response.GetError().Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
