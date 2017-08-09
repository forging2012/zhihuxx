/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License
*/
package zhihuxx

import "fmt"

var (
	// 谁关注
	followersrurl = "https://www.zhihu.com/api/v4/members/%s/followers?"

	// 关注谁
	followeesurl = "https://www.zhihu.com/api/v4/members/%s/followees?include="
	fparm        = "data[*].answer_count,articles_count,gender,follower_count,is_followed,is_following,badge[?(type=best_answerer)].topics"
)

type FollowData struct {
	Page PageInfo           `json:"paging"`
	Data []FollowerDataInfo `json:"data"`
}

/*
"is_followed": false,
"avatar_url_template": "https://pic3.zhimg.com/f9c55d14d855e0bf44511bd9c0f73aae_{size}.jpg",
"user_type": "people",
"answer_count": 13,
"is_following": false,
"headline": "环境工程专业学生",
"url_token": "li-bin-63-8-70",
"id": "98a410425c393a2a64c35bdca8ad8f27",
"articles_count": 0,
"type": "people",
"name": "李斌",
"url": "http://www.zhihu.com/api/v4/people/98a410425c393a2a64c35bdca8ad8f27",
"gender": -1,
"is_advertiser": false,
"avatar_url": "https://pic3.zhimg.com/f9c55d14d855e0bf44511bd9c0f73aae_is.jpg",
"is_org": false,
"follower_count": 12,
"badge": []
*/
type FollowerDataInfo struct {
}

func Followees(token string) string {
	return fmt.Sprintf(followeesurl, token) + fparm + "&limit=%d&offset=%d"
}

func Followers(token string) string {
	return fmt.Sprintf(followeesurl, token) + fparm + "&limit=%d&offset=%d"
}
