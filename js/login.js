
const userEmail = document.getElementById("user-email")
const userPassword = document.getElementById("user-password")
const userLoginButton = document.getElementById("user-login-button")
const content = document.getElementById("content")

if (sessionStorage.getItem("user") === null || sessionStorage.getItem("user") === "") {
    content.style.visibility = "visible"
    console.log("session is not logged in")
} else {
    console.log("session is logged in")
    verifyUser();
}

userLoginButton.addEventListener("click", () => {
    verifyUser();
})