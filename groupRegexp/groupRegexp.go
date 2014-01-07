package groupRegexp

import "regexp"

type GroupRegexp struct {
	regexTxt string
	text string
	regex *regexp.Regexp
	group2match map[string]string
}

func NewGroupRegexp(regexTxt, text string) *GroupRegexp {
	return new(GroupRegexp).MustCompile(regexTxt, text)
}

func (mr *GroupRegexp) MustCompile(regexTxt, text string) *GroupRegexp {
	mr.regexTxt = regexTxt
	mr.regex = regexp.MustCompile(mr.regexTxt)
	mr.text = text

	return mr
}

func (mr *GroupRegexp) GetGroup(group string) string {
	if mr.regex == nil {
		panic("Must compile first")
	}
	if mr.group2match != nil && len(mr.group2match) != 0 {
		return mr.group2match[group]
	}

    groups := mr.regex.SubexpNames()
    submatches := mr.regex.FindAllStringSubmatch(mr.text, -1)[0]
    mr.group2match = map[string]string {}
    for i, val := range submatches {
        mr.group2match[groups[i]] = val
    }

    return mr.group2match[group]
}
