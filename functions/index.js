const functions = require('firebase-functions');
const admin = require('firebase-admin');
admin.initializeApp();

exports.createProfile = functions.auth.user().onCreate((user) => {
  return admin.firestore().collection('profiles').doc(user.uid).set({
    email: user.email,
  });
});

exports.deleteProfile = functions.auth.user().onDelete(async (user) => {
  return admin.firestore().collection('profiles').doc(user.uid).delete();
});
