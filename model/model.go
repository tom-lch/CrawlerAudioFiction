package model

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
)

type JingTingData struct {
	ID         string `sql:"id"`
	Title      string `sql:"title"`       // 小说名称
	Class      string `sql:"class"`       // 类别
	Author     string `sql:"author"`      // 作者
	Announcer  string `sql:"annonucer'`   // 播音
	Hits       string `sql:"hists"`       // 人气
	State      string `sql:"state"`       // 状态
	UpdateTime string `sql:"update_time"` // 时间
	Intro      string `sql:"intro"`       // 简介
	Url        string `sql:"url"`
	Host       string `sql:"host"`
	Plists     string `sql:"plist"` // 获取信息的列表
	Plist      []string
}

func (jt *JingTingData) SetID() {
	if len(jt.ID) > 0 {
		return
	}
	info := fmt.Sprintf("%s/%s/%s", jt.Title, jt.Announcer, jt.Url)

	h := fnv.New32a()
	h.Write([]byte(info))
	val := h.Sum32()
	jt.ID = strconv.Itoa(1000+rand.Intn(1000)) + strconv.Itoa(int(val))
}

type BiliBiliSearchPageAll struct {
	Data BiliBiliSearchPageAllData `json:"data"`
}

type BiliBiliSearchPageAllData struct {
	NumPages   int `json:"numPages"`
	NumResults int `json:"numPages"`
	Page       int `json:"numPages"`
	PageSize   int `json:"numPages"`
}

type BiliBilPageInfo struct {
	Data BiliBilEveryPageInfo `json:"data"`
}

type BiliBilEveryPageInfo struct {
	Result []VideoInfo `json:"result"`
}

type VideoInfo struct {
	ID           string
	AID          int      `json:"aid"`
	Arcrank      string   `json:"arcrank"`
	Arcurl       string   `json:"arcurl"`
	Author       string   `json:"author"`
	Badgepay     bool     `json:"badgepay"`
	Bvid         string   `json:"bvid"`
	Description  string   `json:"description"`
	Duration     string   `json:"duration"`
	Favorites    int      `json:"favorites"`
	HitColumns   []string `json:"hit_columns"`
	VID          int      `json:"id"`
	IsPay        int      `json:"is_pay"`
	IsUnionVideo int      `json:"is_union_video"`
	MID          int      `json:"mid"`
	NewRecTags   []string `json:"new_rec_tags"`
	Pic          string   `json:"pic"`
	Play         int      `json:"play"`
	Pubdate      int      `json:"pubdate"`
	RankScore    int      `json:"rank_score"`
	Review       int      `json:"review"`
	Senddata     int      `json:"senddate"`
	Tag          string   `json:"tag"`
	Title        string   `json:"title"`
	Type         string   `json:"type"`
	Typeid       string   `json:"typeid"`
	Typename     string   `json:"typename"`
	VideoReview  int      `json:"video_review"`
	ViewType     string   `json:"view_type"`
}

func (vi *VideoInfo) SetID() {
	if len(vi.ID) > 0 {
		return
	}
	info := fmt.Sprintf("%s/%s/%s/%d/%d", vi.Title, vi.Bvid, vi.Author, vi.AID, vi.VID)

	h := fnv.New32a()
	h.Write([]byte(info))
	val := h.Sum32()
	vi.ID = strconv.Itoa(2000+rand.Intn(1000)) + strconv.Itoa(int(val))
}

type XimalayaPageInfo struct {
	Data XimalayaData `json:"data"`
}

type XimalayaData struct {
	Albums []Album `json:"albums"`
}

type Album struct {
	ID             string
	AlbumId        int    `json:"albumId"`
	AlbumSubscript int    `json:"albumSubscript"`
	AnchorName     string `json:"anchorName"`
	CoverPath      string `json:"coverPath"`
	IsFinished     int    `json:"isFinished"`
	IsPaid         bool   `json:"isPaid"`
	Link           string `json:"link"`
	PlayCount      int    `json:"playCount"`
	Title          string `json:"title"`
	TrackCount     int    `json:"trackCount"`
	Uid            int    `json:"uid"`
	VipType        int    `json:"vipType"`
}

func (al *Album) SetID() {
	if len(al.ID) > 0 {
		return
	}
	info := fmt.Sprintf("%s/%s/%s/%d/%d", al.Title, al.Link, al.AnchorName, al.Uid, al.AlbumId)
	h := fnv.New32a()
	h.Write([]byte(info))
	val := h.Sum32()
	al.ID = strconv.Itoa(3000+rand.Intn(1000)) + strconv.Itoa(int(val))
}
