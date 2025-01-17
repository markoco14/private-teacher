let faqItems = document.querySelectorAll('dt');
        faqItems.forEach(item => {
            item.addEventListener('click', async () => {
            item.nextElementSibling.classList.toggle('hidden')
            })
        })