// Get popup elements
const privacyPopup = document.getElementById("privacy-popup");
const contactPopup = document.getElementById("contact-popup");

// Get links and close buttons
const privacyLink = document.getElementById("privacy-link");
const contactLink = document.getElementById("contact-link");
const closePrivacy = document.getElementById("close-privacy");
const closeContact = document.getElementById("close-contact");

// Show the privacy popup
privacyLink.addEventListener("click", function (e) {
  e.preventDefault();
  privacyPopup.style.display = "flex";
});

// Show the contact popup
contactLink.addEventListener("click", function (e) {
  e.preventDefault();
  contactPopup.style.display = "flex";
});

// Close the privacy popup
closePrivacy.addEventListener("click", function () {
  privacyPopup.style.display = "none";
});

// Close the contact popup
closeContact.addEventListener("click", function () {
  contactPopup.style.display = "none";
});

// Close popups when clicking outside the content area
window.addEventListener("click", function (e) {
  if (e.target === privacyPopup) {
    privacyPopup.style.display = "none";
  }
  if (e.target === contactPopup) {
    contactPopup.style.display = "none";
  }
});
