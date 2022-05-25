function qs(elm) {
  return document.querySelector(elm);
}

const cReg = document.querySelector(".btn-register");
const cLog = document.querySelector(".btn-log-in");

cReg.addEventListener("click", () => {
  document.querySelector(".container-login").style.display = "none";
  document.querySelector(".container-register").style.display = "flex";
  return;
});

cLog.addEventListener("click", () => {
  document.querySelector(".container-register").style.display = "none";
  document.querySelector(".container-login").style.display = "flex";
  return;
});

const xml = new XMLHttpRequest();

const btnLogin = qs(".btn-login");
const email = qs("#input-email");
const password = qs("#input-password");
btnLogin.addEventListener("click", function (e) {
  /* for sub domain */
  // xml.withCredentials = true;
  xml.open("POST", `${window.origin}/login`, true);
  xml.setRequestHeader("content-type", "application/json");
  // email: "rikianfaisal@gmail.com",
  // password: "54n94t_r4h4514...",
  xml.send(
    JSON.stringify({
      email: `${email.value}`,
      password: `${password.value}`,
    })
  );
  return;
});

const btnSignUp = qs(".btn-signup");
const rFullName = qs("#r-fullName");
const rPhone = qs("#r-phone");
const rEmail = qs("#r-email");
const rStatus = qs("#r-status");
const rPassword1 = qs("#r-password1");
const rPassword2 = qs("#r-password2");
const status = qs("#r-fullName");

btnSignUp.addEventListener("click", function (e) {
  xml.open("POST", `${window.origin}/register`, true);
  xml.setRequestHeader("content-type", "application/json");
  // password: "54n94t_r4h4514...",
  if (
    !rEmail.value.match(/^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/) ||
    rEmail.value.length > 32 ||
    rPassword1.value !== rPassword2.value ||
    rPassword1.value.length > 32
  )
    return alert("Invalid data");
  const data_user = JSON.stringify({
    user_email: rEmail.value,
    user_name: rFullName.value,
    user_password: rPassword1.value,
    user_phone: rPhone.value,
    user_status: rStatus.value,
  });
  xml.send(data_user);
  return;
});

xml.onload = function (e) {
  if (xml.status != 200) {
    return alert(JSON.parse(xml.response)["message"]);
  }

  alert(JSON.parse(xml.response)["message"]);
  window.location.reload();
  return;
};
