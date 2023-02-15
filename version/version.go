package version

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	head   int64
	middle int64
	tail   int64
}

func NewVersion(head int64, middle int64, tail int64) *Version {
	return &Version{
		head:   head,
		middle: middle,
		tail:   tail,
	}
}

func (v *Version) ToString() string {
	return "v" + strconv.FormatInt(v.head, 10) + "." + strconv.FormatInt(v.middle, 10) + "." + strconv.FormatInt(v.tail, 10)
}

func StrToVersion(vstr string) (*Version, error) {
	vstr = strings.TrimSpace(vstr)
	vstr = strings.ToLower(vstr)
	vstr = strings.TrimPrefix(vstr, "v")

	v_array := strings.Split(vstr, ".")

	if len(v_array) != 3 {
		return nil, errors.New("version format error,vstr:" + vstr)
	}

	head_, err := strconv.ParseInt(v_array[0], 10, 64)
	if err != nil {
		return nil, errors.New("version head format error,vstr:" + vstr)
	}

	mid_, err := strconv.ParseInt(v_array[1], 10, 64)
	if err != nil {
		return nil, errors.New("version middle format error,vstr:" + vstr)
	}

	tail_, err := strconv.ParseInt(v_array[2], 10, 64)
	if err != nil {
		return nil, errors.New("version tail format error,vstr:" + vstr)
	}

	return &Version{
		head:   head_,
		middle: mid_,
		tail:   tail_,
	}, nil

}

// a>b return 1 , a==b return 0 ,  a<b  return -1
func VersionCompare(a *Version, b *Version) int {

	//head
	if a.head < b.head {
		return -1
	} else if a.head > b.head {
		return 1
	} else {

		//middle
		if a.middle < b.middle {
			return -1
		} else if a.middle > b.middle {
			return 1
		} else {

			//tail
			if a.tail < b.tail {
				return -1
			} else if a.tail > b.tail {
				return 1
			} else {
				return 0
			}

		}
	}
}
