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


document.addEventListener('mousemove', function (e) {
    // Function to create a single sparkle
    const createSparkle = (x, y) => {
      const sparkle = document.createElement('div');
      sparkle.className = 'sparkle';
      sparkle.style.left = `${x}px`;
      sparkle.style.top = `${y}px`;
  
      // Slightly randomize size and position for a subtle effect
      sparkle.style.transform = `translate(${Math.random() * 6 - 3}px, ${
        Math.random() * 6 - 3
      }px) scale(${Math.random() * 0.5 + 0.5})`;
  
      document.body.appendChild(sparkle);
  
      // Remove the sparkle after animation ends
      sparkle.addEventListener('animationend', function () {
        sparkle.remove();
      });
    };
  
    // Generate fewer sparkles near the cursor
    for (let i = 0; i < 3; i++) { // Reduced to 3 sparkles
      createSparkle(e.pageX, e.pageY);
    }
  });
  