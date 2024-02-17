import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

// TODO: Replace the following with your app's Firebase project configuration
// See: https://firebase.google.com/docs/web/learn-more#config-object
// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyCakBiq0gXKvHE5bD0MYO4Jx3VC5A6oSq8",
    authDomain: "intro-8f3f7.firebaseapp.com",
    projectId: "intro-8f3f7",
    storageBucket: "intro-8f3f7.appspot.com",
    messagingSenderId: "932334191987",
    appId: "1:932334191987:web:bc284b4e2865b7e3e03106",
    measurementId: "G-NB5DFNTF0J"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

function Login() {
    console.log("Login");

    const email = "kenya200214@gmail.com"
    const password = "12345678"

    const auth = getAuth(app);
    signInWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
        // Signed in 
        const user = userCredential.user;
        // ...

        console.log("Login Success! user=", user);
    })
    .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;

        console.log("Login Failed! errorCode=", errorCode, " errorMessage=", errorMessage);
    });
}

// webpackを使用しているため、buldle.jsファイルでLogin関数名は別の物へ置き換わる
// https://qiita.com/Anaakikutsushit/items/18fd24ac77cf379a422c
document.getElementById("login").addEventListener("click", Login);
