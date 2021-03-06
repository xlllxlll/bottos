// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  provide a interface error definition for all modules
 * @Author: zl
 * @Date: 2018-4-25
 * @Last Modified by:
 * @Last Modified time:
*/
package error

import (
	"io/ioutil"
	"encoding/json"
)

type ErrorCode struct {
	Code    int64 `json:"code"`
	Lv      string  `json:"lv"`
	Msg     struct {
		Cn string `json:"cn"`
		En string `json:"en"`
	} `json:"msg"`
	Details string  `json:"details"`
}

func GetErrorInfo(code int64, serviceName string) ErrorCode {
	d := GetAllErrorInfos(serviceName)
	for _, v := range d {
		if code == v.Code {
			v.Code = getServerId(serviceName)*10000 + code
			return v
		}
	}
	return ErrorCode{}
}

func GetAllErrorInfos(serviceName string) []ErrorCode {
	fr, err := ioutil.ReadFile("./"+serviceName+"-ErrorCode.json")
	if err != nil {
		panic(err)
	}

	var d []ErrorCode
	err = json.Unmarshal(fr, &d)
	if err != nil {
		panic(err)
	}
	return d
}


func getServerId(serviceName string) int64 {
	//TODO
	return 0
}

