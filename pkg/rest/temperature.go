// Copyright 2024 Fabian `xx4h` Sylvester
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

package rest

import "fmt"

func (h *Hass) TemperatureSet(obj string, temp float64) (string, string, string, error) {
	svc := "set_temperature"
	sub, obj, err := h.entityArgHandler([]string{obj}, svc)
	if err != nil {
		return "", "", "", err
	}

	payload := map[string]any{
		"entity_id":   fmt.Sprintf("%s.%s", sub, obj),
		"temperature": fmt.Sprintf("%.1f", temp),
	}
	res, err := h.api("POST", fmt.Sprintf("/services/%s/%s", sub, svc), payload)
	if err != nil {
		return "", "", "", err
	}

	if err := h.getResult(res); err != nil {
		return "", "", "", err
	}
	return obj, fmt.Sprintf("temperature set to %.1f", temp), sub, nil

}
