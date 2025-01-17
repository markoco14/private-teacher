document.addEventListener('formSuccess', (event) => {
    const alert = document.getElementById('alert')
    const message = event.detail?.message || 'Thank you for your message. We will get back to you soon.'
    
    alert.innerText = message
    alert.classList.remove('-translate-y-8', 'opacity-0')

    setTimeout(() => {
        alert.classList.add('-translate-y-8', 'opacity-0')
    }, 3000)
})