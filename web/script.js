console.log("Its a live!");

const button = document.querySelector('.btn'),
      emailInput = document.querySelector('.email'),
      passwordInput = document.querySelector('.password'),
      headerBtn = document.querySelector('.header_btn'),
      registrationPopup = document.querySelector('.registration_wrap')

const URL = 'http://localhost:8080/users'

headerBtn.addEventListener('click', () => {
  registrationPopup.classList.toggle('non_visible')
})

// handle of registration button
button.addEventListener('click', () => {
  fetch(URL, {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: emailInput.value,
      password: passwordInput.value
    })
  })

  console.log(emailInput.value);
  console.log(passwordInput.value);
})

