{{ define "contact" }}
<section id="contact" class="bg-gradient-to-r from-purple-700/80 to-purple-800/80 text-white py-24 md:py-24">
    <div class="hide max-w-[500px] mx-auto bg-white text-black px-6 py-6 sm:rounded-2xl">
        <h2 class="text-2xl font-semibold text-purple-900 mb-2">
            {{ if eq .Lang "en" }}
                Contact Me To Arrange Your First Class
            {{ else }}
                聯絡我安排第一堂課
            {{ end }}
        </h2>
        <p>
            {{ if eq .Lang "en" }}
                Scan the Line QR code or fill out the email form below. I am available between 8 am and 8 pm everyday to respond to your inquiries. I will do my best to reply to you within 2 hours of messaging me.
            {{ else }}
                掃描Line QR碼或填寫下面的電子郵件表格。我每天上午8點至晚上8點都可以回答您的問題。我將盡力在您發送消息後的2小時內回复您。
            {{ end }}
        </p>
        <div  class="my-4 text-center text-purple-900 font-semibold text-xl">
            <h3>Line</h3>
            <hr />
        </div>
        <img 
            id="line" 
            loading="lazy" 
            src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/tr-mark-qr.jpg" 
            class="w-full max-w-[200px] mx-auto"/>
        <div  class="my-4 text-center text-purple-900 font-semibold text-xl">
            <h3>
            {{ if eq .Lang "en" }}
                Email
            {{ else }}
                電子郵件
            {{ end }}
            </h3>
            <hr />
        </div>
        {{ template "form" . }}
    </div>
</section>
{{ end }}