
const button = document.querySelector('.btn'),
      emailInput = document.querySelector('.email'),
      passwordInput = document.querySelector('.password'),
      headerBtn = document.querySelector('.header_btn'),
      registrationPopup = document.querySelector('.registration_wrap'),
      usersBtn = document.querySelector('.users_btn'),
      userList = document.querySelector('.userlist')

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
})

// handle of all users button
usersBtn.addEventListener('click', () => {  // TODO add errors handler
  fetch(URL)
  .then(data => data.json())
  .then(output => {
    for (item in output) {
      let li = document.createElement("li")
      let node = document.createTextNode(output[item])
      li.appendChild(node)
      userList.appendChild(li)
    }
  })
})

