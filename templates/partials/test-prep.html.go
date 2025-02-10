{{ define "test-prep" }}
<section id="test-prep" class="min-h-[50vh] grid place-content-center py-32 bg-gray-50">
    <div class="hide max-w-[500px] md:max-w-[800px] mx-auto px-4 sm:px-0">
        <h2 class="font-bold text-deep text-4xl text-balance md:text-center mb-8">
        {{ if eq .Lang "en" }}
            Test preparation classes for IELTS, GEPT, and more.
        {{ else }}
            IELTS、GEPT等考試準備課程。
        {{ end }}
        </h2>
        <div class="mx-auto grid grid-cols-1 md:grid-cols-3 gap-4 mb-16">
            <div class="md:shadow-xl rounded-xl md:text-center bg-white">
                <h3 class="py-2 px-4 md:text-center text-xl text-deep font-semibold">IELTS</h3>
                <div class="stacked overflow-y-hidden w-full h-[250px] md:h-[150px] bg-purple-500">
                    <img 
                        height="200"
                        width="200"
                        loading="lazy" 
                        src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-prep-student-with-book.jpg" 
                        alt="A girl looking up from a book."
                        class="w-full h-full object-cover"/>
                    <div class="w-full h-[250px] md:h-[150px] bg-purple-300/30"></div>
                </div>
                <p class="py-2 px-4">
                    {{ if eq .Lang "en" }}
                        Study and work in English speaking countries.
                    {{ else }}
                        學習和在英語國家工作。
                    {{ end }}
                </p>
            </div>
            <div class="md:shadow-xl rounded-xl md:text-center bg-white">
                <h3 class="py-2 px-4 md:text-center text-xl text-deep font-semibold">GEPT</h3>
                <div class="stacked overflow-y-hidden w-full h-[250px] md:h-[150px] bg-purple-500">
                    <img 
                        height="200"
                        width="200"
                        loading="lazy" 
                        src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-prep-students-classroom.jpg" 
                        alt="A girl looking up from a book."
                        class="w-full h-full object-cover"/>
                    <div class="w-full h-[250px] md:h-[150px] bg-purple-300/30"></div>
                </div>
                <p class="py-2 px-4">
                    {{ if eq .Lang "en" }}
                        Elevate your education and career possibilities.
                    {{ else }}
                        提升您的教育和職業機會。
                    {{ end }}
                </p>
            </div>
            <div class="md:shadow-xl rounded-xl md:text-center bg-white">
                <h3 class="py-2 px-4 md:text-center text-xl text-deep font-semibold">
                    {{ if eq .Lang "en" }}
                        Other
                    {{ else }}
                        其他
                    {{ end }}
                </h3>
                <div class="stacked overflow-y-hidden w-full h-[250px] md:h-[150px] bg-purple-500">
                    <img 
                        height="200"
                        width="200"
                        loading="lazy" 
                        src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-prep-write-test.jpg" 
                        alt="A girl looking up from a book."
                        class="w-full h-full object-cover"/>
                    <div class="w-full h-[250px] md:h-[150px] bg-purple-300/30"></div>
                </div>
                <p class="py-2 px-4">
                {{ if eq .Lang "en" }}
                    Take any language test with confidence.
                {{ else }}
                    有信心參加任何語言測試。
                {{ end }}
                </p>
            </div>
        </div>
        <div class="flex justify-center gap-4">
            <a href="#contact" class="inline-block bg-purple-800 px-8 py-2 text-white rounded-md hover:bg-purple-700 active:bg-purple-900 active:scale-95 duration-200 ease-in-out">
            {{ if eq .Lang "en" }}
                Book Now
            {{ else }}
                立即預約
            {{ end }}
            </a>
            {{/* <a href="#test-prep" class="inline-block text-lg bg-purple-300 text-purple-800 rounded-md px-4 py-1 hover:bg-purple-200 active:bg-purple-400 active:text-purple-900 active:scale-95 duration-200 ease-in-out">
            {{ if eq .Lang "en" }}
                Learn More
            {{ else }}
                了解更多
            {{ end }}
            </a> */}}
        </div>
    </div>
</section>
{{ end }}