/**
 * Import function triggers from their respective submodules:
 *
 * const {onCall} = require("firebase-functions/v2/https");
 * const {onDocumentWritten} = require("firebase-functions/v2/firestore");
 *
 * See a full list of supported triggers at https://firebase.google.com/docs/functions
 */

const { onRequest } = require("firebase-functions/v2/https");
const { user } = require("firebase-functions/v1/auth");
const { initializeApp } = require("firebase-admin/app");
const admin = require("firebase-admin");
const logger = require("firebase-functions/logger");

// Create and deploy your first functions
// https://firebase.google.com/docs/functions/get-started

// exports.helloWorld = onRequest((request, response) => {
//   logger.info("Hello logs!", {structuredData: true});
//   response.send("Hello from Firebase!");
// });

// Create a user information document in Firestore after creating a new user
exports.createUserInfoDocument = user().onCreate((user) => {
    initializeApp();

    const user_info = {
        id: user.uid,
        gender: "",
        health_information: {
            birthday: "",
            weight: "",
            height: "",
            bmi: "",
            exercise_frequency: "",
        },
        exercise_frequency: [],
    };

    return admin
        .firestore()
        .collection("user_info")
        .doc(user.uid)
        .set(user_info);
});
