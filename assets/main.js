var url = "ws://" + window.location.host + "/sampleapp/ws";
var ws = new WebSocket(url);
var name = "User-" + Math.floor(Math.random() * 1000);
var chat = document.getElementById("chat");
document.getElementById("user").placeholder = name;

var text = document.getElementById("text");
var now = function () {
    return new Date().toLocaleString();
};

ws.onmessage = function (msg) {
    if(msg.data.startsWith("R:")){
        var nums = msg.data.slice(2);
        number.innerText = nums;
    }
    else{
        var line =  now() + " : " + msg.data + "\n";
        chat.innerText += line;
    }
};

text.onkeydown = function (e) {
    if (e.keyCode === 13 && text.value !== "") {
        var nowname = document.getElementById("user").value;
        console.log(nowname);
        ws.send("[" + ((nowname != "") ? nowname : name) + "] > " + text.value);
        text.value = "";
    }
};

function roll(){
    ws.send("Roll");
}
