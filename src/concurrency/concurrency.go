package main

import "fmt"

/* Concurrency la kha nang 1 chuong trinh co the dieu phoi nhieu tac vu trong cung mot khoang thoi gian
*  va trong qua trinh dieu phoi chi cho phep 1 tac vu chay trong 1 thoi diem
 */

/* Parallelism là khả năng 1 chương trình có thể thực thi 2 hoặc nhiều tasks trong cùng một thời điểm
* với điều kiện CPU có từ 2cores trở lên
 */

/* Do not communicate by sharing memory; instead, share memory by communicating
* Khi giao tiếp giữa các go routines thì sẽ k trực tiếp chia sẻ bộ nhớ, mà sẽ thông qua channel
 */

