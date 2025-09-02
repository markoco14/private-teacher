const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        if (entry.isIntersecting) {
            entry.target.classList.add('show')
        } 
    })
})

const hiddenElements = document.querySelectorAll('.hide')
hiddenElements.forEach((element) => {
    observer.observe(element)
})