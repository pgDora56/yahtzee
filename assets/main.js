var url = "ws://" + window.location.host + "/sampleapp/ws";
var ws = new WebSocket(url);
var name = "User-" + Math.floor(Math.random() * 1000);
var log = document.getElementById("log");
document.getElementById("user").placeholder = name;

var text = document.getElementById("text");
var now = function () {
    return new Date().toLocaleString();
};

ws.onmessage = function (msg) {
    if(msg.data.startsWith("R:")){
        var nums = msg.data.slice(3).split(" ");
        console.log(nums);
        document.getElementById("dicel0").innerText = nums[0];
        document.getElementById("dicel1").innerText = nums[1];
        document.getElementById("dicel2").innerText = nums[2];
        document.getElementById("dicel3").innerText = nums[3];
        document.getElementById("dicel4").innerText = nums[4];
    }
    else{
        var line =  now() + " : " + msg.data + "\n";
        log.innerText += line;
    }
};

// text.onkeydown = function (e) {
//     if (e.keyCode === 13 && text.value !== "") {
//         var nowname = document.getElementById("user").value;
//         console.log(nowname);
//         ws.send("[" + ((nowname != "") ? nowname : name) + "] > " + text.value);
//         text.value = "";
//     }
// };

function roll(){
    console.log("rolling");
    var dlis = "";
    if(document.getElementById("dice0").checked){
        dlis += "0"
    } 
    if(document.getElementById("dice1").checked){
        dlis += "1"
    } 
    if(document.getElementById("dice2").checked){
        dlis += "2"
    }
    if(document.getElementById("dice3").checked){
        dlis += "3"
    }
    if(document.getElementById("dice4").checked){
        dlis += "4"
    }
    if(dlis.length == 0) ws.send("Roll");
    else {
        ws.send("Reroll:" + dlis)
    }
}
