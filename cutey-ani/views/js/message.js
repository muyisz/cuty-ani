$(document).ready(function () {
    function changeDateFormat(cellval) {
        let dateVal = cellval + '';
        if (cellval != null) {
            let reg = new RegExp('.\\d{3}\\+\\d{4}$');
            let date = new Date(dateVal.replace(reg, '').replace('T', ' '));
            let month =
                date.getMonth() + 1 < 10
                    ? '0' + (date.getMonth() + 1)
                    : date.getMonth() + 1;
            let day = date.getDate() < 10 ? '0' + date.getDate() : date.getDate();
            return date.getFullYear() + '-' + month + '-' + day;
        }
    }
    function GetChatRoom() {
        $.ajax({
            url: '/getRoom',
            type: 'get',
            dataType: 'json',
            async: false,
            cache: false,
            processData: false,
            contentType: false,
            success: function (data) {
                console.log(data)
                if (data.pass) {
                    head = ""
                    tail = "<br>"
                    k = ""
                    for (i = 0; i < data.msg.length; i++) {
                        k = k + head + "<h5>" + data.msg[i].From + "</h5>" + "<h3>" + data.msg[i].Content + "</h3>" + "<h7>" + changeDateFormat(data.msg[i].Time) + "</h7>" + tail
                    }
                    document.getElementById("chat_box").innerHTML = k;
                }
                else {
                    alert("聊天内容获取失败1！")
                }
            },
            error: function () {
                alert("聊天内容获取失败2！")
            }
        })
    }
    GetChatRoom()
    $(function k() {
        setTimeout(GetChatRoom(),10000000);
        k()
    });
    $('#send').click(function () {
        const msg = document.getElementById('msg').value;
        if (msg != "") {
            var formdata = new FormData()
            formdata.append("msg", msg)
            $.ajax({
                url: '/postRoom',
                type: 'post',
                dataType: 'json',
                data: formdata,
                processData: false,
                contentType: false,
                success: function (data) {
                    if (data.pass) {
                        GetChatRoom()
                    } else {
                        alert("发送失败1")
                    }
                },
                error: function () {
                    alert("发送失败2")
                }
            })
        }
    })
})
