package main

// import (
// 	"strings"
// 	"testing"
// )

// func solution(id_list []string, report []string, k int) []int {
//     // 신고한 유저 -> 신고당한 유저
//     alertUserMap := make(map[string][]string)
//     // 신고당한 횟수
//     alertUserCountMap := make(map[string]int)
//     // 이메일 접수 유저
//     emailSendMap :=make(map[string]int)
//     // 이메일 정지 유저
//     emailRejectUser := make([]string, len(report))

//     for _, r := range report {
//         strs :=strings.Split(r, " ")

//         reportingUser := strs[0]
//         reportedUser := strs[1]

//         val, _ := alertUserMap[reportingUser]
//         count, _ := alertUserCountMap[reportedUser]

//         val = append(val,reportedUser)
//         alertUserMap[reportingUser] = val

//         alertUserCountMap[reportedUser] = count+1
//     }

//     for key,v := range alertUserCountMap {
//         if v == k {
//             emailRejectUser = append(emailRejectUser, key)
//         }
//     }

//     for _, user := range emailRejectUser {

//         if user == "" {
//             continue
//         }

//         for k,v := range alertUserMap {
//             userStr := strings.Join(v,"")

//             if isContainer := strings.Contains(userStr, user); isContainer {
//                 val, _ := emailSendMap[k]
//                 emailSendMap[k] = val+1
//             }
//         }
//     }

//     result := make([]int, len(id_list))
//     for index, id := range id_list {
//         result[index] = emailSendMap[id]
//     }

//     return result
// }

// // test code
// func Test_solution_1(t *testing.T) {
// 	id_list := []string{"muzi", "frodo", "apeach", "neo"}
// 	report := []string{"muzi frodo","apeach frodo","frodo neo","muzi neo","apeach muzi"}
// 	k := 2

// 	result := solution(id_list, report, k)

// 	for index, item := range []int{2,1,1,0} {
// 		if item != result[index]{
// 			t.Error("not equlas")
// 		}
// 	}
// }

// func Test_solution_2(t *testing.T) {
// 	id_list := []string{"con", "ryan"}
// 	report := []string{"ryan con", "ryan con", "ryan con", "ryan con"}
// 	k := 3

// 	result := solution(id_list, report, k)

// 	for index, item := range []int{0,0} {
// 		if item != result[index]{
// 			t.Error("not equlas")
// 		}
// 	}
// }