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
  var img = document.getElementById("confirmEyeImg");
  if (x.type === "password") {
    x.type = "text";
    img.src = "../../../static/img/pw-see.svg"
  } else {
    x.type = "password";
    img.src = "../../../static/img/pw-hide.svg"
  }
}

const button = document.getElementById("toggleEye");
const confirmButton = document.getElementById("confirmToggleEye");
button.addEventListener("click", togglePWvisible);
confirmButton.addEventListener("click", togglePWconfirmVisible);
