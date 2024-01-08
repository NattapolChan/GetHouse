// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { getFirestore } from "firebase/firestore";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
	apiKey: process.env.REACT_FIREBASE_API_KEY,
	authDomain: process.env.REACT_FIREBASE_AUTH_DOMAIN,
	projectId: process.env.REACT_FIREBASE_PROJECT_ID,
	storageBucket: process.env.REACT_FIREBASE_STORAGE_BUCKET,
	messagingSenderId: process.env.REACT_FIREBASE_MESSAGING_SENDER_ID,
	appId: process.env.REACT_FIREBASE_APP_ID,
//  apiKey: "AIzaSyBqP3Op_UJkamcunIRATatmZ2AfrmRj38I",
//  authDomain: "react-frontend-f3529.firebaseapp.com",
//  projectId: "react-frontend-f3529",
//  storageBucket: "react-frontend-f3529.appspot.com",
//  messagingSenderId: "561340384167",
//  appId: "1:561340384167:web:cb5a498c372239c4e12c82",
};

// Initialize Firebase
export const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const db = getFirestore(app);
