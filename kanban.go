package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type KanbanManager struct {
	Kanbans []*Kanban

	mode                  string
	currentImpotingKanban *Kanban
}

type Kanban struct {
	Title          string
	TodoList       []*Issue
	InProgressList []*Issue
	DoneList       []*Issue
}

type Issue struct {
	Name  string
	Depth int
	Check bool
	Tags  []string
}

func (k *KanbanManager) importTodomdFile(filename string) ([]*Kanban, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	k.mode = "start"
	for scanner.Scan() {
		k.runLine(scanner.Text())
	}
	return k.Kanbans, nil
}

func (k *KanbanManager) runLine(line string) {
	if strings.TrimSpace(line) == "" {
		return
	}

	if strings.Index(line, "# ") == 0 {
		k.currentImpotingKanban = &Kanban{
			Title: line[2:],
		}
		k.Kanbans = append(k.Kanbans, k.currentImpotingKanban)
		return
	}

	if strings.Index(line, "### TODO") == 0 {
		k.mode = "todo"
		return
	}

	if strings.Index(line, "### In Progress") == 0 {
		k.mode = "in_progress"
		return
	}

	if strings.Index(line, "### Done") == 0 {
		k.mode = "done"
		return
	}

	if k.mode == "" {
		return
	}

	// TODO: Check Vaild

	depth := strings.Index(line, "-") / 2
	issueLine := strings.TrimSpace(line)
	issueLine = issueLine[2:]
	check := false
	if strings.Index(issueLine, "[ ]") == 0 {
	} else if strings.Index(issueLine, "[-]") == 0 {
	} else if strings.Index(issueLine, "[x]") == 0 {
		check = true
	}
	issueLine = issueLine[4:]

	var tags []string
	issueLine, tags = extractHashtags(issueLine)

	issue := &Issue{
		Name:  issueLine,
		Depth: depth,
		Check: check,
		Tags:  tags,
	}
	switch k.mode {
	case "todo":
		k.currentImpotingKanban.TodoList = append(k.currentImpotingKanban.TodoList, issue)
	case "in_progress":
		k.currentImpotingKanban.InProgressList = append(k.currentImpotingKanban.InProgressList, issue)
	}
}

// extractHashtags 이슈 내용을 태그와 분리 시켜주는 코드. by ChatGPT
func extractHashtags(input string) (string, []string) {
	// 해시태그를 찾기 위한 정규 표현식 패턴
	regexPattern := `#([0-9A-Za-z]+)`
	re := regexp.MustCompile(regexPattern)

	// 정규 표현식에 매치되는 모든 문자열 찾기
	matches := re.FindAllString(input, -1)

	hashtags := make([]string, 0, len(matches))
	for _, match := range matches {
		hashtags = append(hashtags, match[1:]) // # 기호를 제외한 문자열 저장
	}

	// 해시태그를 제거한 문자열 생성
	noHashtagsText := re.ReplaceAllString(input, "")
	noHashtagsText = strings.TrimSpace(noHashtagsText)

	return noHashtagsText, hashtags
}
