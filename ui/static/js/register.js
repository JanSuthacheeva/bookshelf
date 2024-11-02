function togglePWvisible() {
  var x = document.getElementById("password");
  var img = document.getElementById("eyeImg");
  if (x.type === "password") {
    x.type = "text";
    img.src = "../../../static/img/pw-see.svg"
  } else {
    x.type = "password";
    img.src = "../../../static/img/pw-hide.svg"
  }
}

function togglePWconfirmVisible() {
    var x = document.getElementById("password_confirm");
    var img = document.getElementById("eyeImg2");
    if (x.type === "password") {
      x.type = "text";
      img.src = "../../../static/img/pw-see.svg"
    } else {
      x.type = "password";
      img.src = "../../../static/img/pw-hide.svg"
    }
}

function addListeners() {
    const button1 = document.getElementById("toggleEye");
    const button2 = document.getElementById("toggleEye2");
    if (button1) {
        button1.addEventListener("click", togglePWvisible);
    }
    if (button2) {
        button2.addEventListener("click", togglePWconfirmVisible);
    }
}

addListeners();

document.addEventListener('htmx:afterOnLoad', addListeners);
