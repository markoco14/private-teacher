<!DOCTYPE html>
<html lang="en">
  <head>
    <title>My App</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="A brief description of your site.">
    <meta name="keywords" content="teaching, lessons, English, tutor, private, online, classes">
    <meta name="author" content="Mark O'Connor">
    <script src="https://unpkg.com/htmx.org@2.0.4" integrity="sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+" crossorigin="anonymous"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        deep: '#020517',
                        deephover: '#0a1971'
                    },
                },
            },
        }
    </script>
    <style>

        .stacked {
            display: grid;
        }

        .stacked > * {
            grid-column: 1 / -1;
            grid-row: 1 / -1;
        }

        .hide {
            opacity: 0;
            filter: blur(2);
            transform: translateY(10%);
            transition: all 1s;
        }

        .show {
            opacity: 1;
            filter: blur(0);
            transform: translateY(0);
        }

        .htmx-indicator {
            display: none;
        }

        .htmx-request .htmx-indicator{
            display:inline;
        }

        .htmx-request.htmx-indicator{
            display:inline;
        }

        .htmx-request .htmx-indicator-hide {
            display: none;
        }

        .htmx-request.htmx-indicator-hide {
            display: none;
        }

        @media(pefers-reduced-motion) {
            .hide {
                transition: none;
            }
        }
    </style>
  </head>
  <body>
    <header class="w-full h-[48px] bg-deep text-white flex items-center z-10">
        <div class="w-full max-w-[1000px] mx-auto px-4 xl:px-0">
            <span>Teacher Mark</span>
        </div>
    </header>
    <main>
        <section class="min-h-[75vh] bg-deep grid place-content-center">
            <div class="hide max-w-[600px] mx-auto px-4 py-4 md:px-0">
                <h1 class="text-5xl text-gray-100 text-pretty mb-4 md:mb-2 font-bold">Achieve confidence, comfort, and clarity in English.</h1>
                <p class="text-lg text-balance text-purple-100 mb-4">Test Preparation and Business English classes to help you achieve your goals. From beginner to advanced, we have classes for your needs and learning style.</p>
                <div class="grid sm:grid-cols-2 text-center gap-4">
                    <a href="/#test-prep" class="inline-block text-lg bg-purple-800 text-white rounded-md hover:bg-purple-700 active:bg-purple-900 active:scale-95 duration-200 ease-in-out">
                        <div class="w-full mx-auto stacked rounded-md">
                            <img 
                                height="200"
                                width="200"
                                src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-hero-write-test.jpg" 
                                class="w-full max-h-[150px] object-cover rounded-md"/>
                            <div class="w-full h-full grid place-content-center bg-purple-900/70 hover:bg-purple-900/80 duration-300 ease-in-out rounded-md">
                                <span class="text-2xl text-white rounded-md duration-300 ease-in-out hover:underline hover:underline-offset-4">Prepare For Test</span>
                            </div>
                        </div>
                    </a>
                    <a href="/#business-english" class="inline-block text-lg bg-purple-300 text-purple-800 rounded-md hover:bg-purple-200 active:bg-purple-400 active:text-purple-900 active:scale-95 duration-200 ease-in-out">
                        <div class="w-full mx-auto stacked rounded-md">
                            <img 
                                height="267"
                                width="400"
                                src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-hero-presentation.jpg" 
                                class="w-full max-h-[150px] object-cover rounded-md"/>
                            <div class="w-full h-full grid place-content-center bg-purple-900/70 hover:bg-purple-900/80 duration-300 ease-in-out rounded-md">
                                <span class="text-2xl text-white rounded-md duration-300 ease-in-out hover:underline hover:underline-offset-4">Improve Business English</span>
                            </div>
                        </div>
                    </a>
                </div>
            </div>
        </section>

        <section id="test-prep" class="min-h-[50vh] grid place-content-center py-32 bg-gray-50">
            <div class="hide max-w-[500px] md:max-w-[800px] mx-auto px-4 sm:px-0">
                <h2 class="font-bold text-deep text-4xl text-balance md:text-center mb-8">Test preparation classes for IELTS, GEPT, and more.</h2>
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
                        <p class="py-2 px-4">Study and work in English speaking countries.</p>
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
                        <p class="py-2 px-4">Elevate your education and career possibilities.</p>
                    </div>
                    <div class="md:shadow-xl rounded-xl md:text-center bg-white">
                        <h3 class="py-2 px-4 md:text-center text-xl text-deep font-semibold">Other</h3>
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
                        <p class="py-2 px-4">Take any language test with confidence.</p>
                    </div>

                </div>
                <div class="flex justify-center gap-4">
                    <a href="#contact" class="inline-block bg-purple-800 px-8 py-2 text-white rounded-md hover:bg-purple-700 active:bg-purple-900 active:scale-95 duration-200 ease-in-out">
                        Book Now
                    </a>
                    <a href="#test-prep" class="inline-block text-lg bg-purple-300 text-purple-800 rounded-md px-4 py-1 hover:bg-purple-200 active:bg-purple-400 active:text-purple-900 active:scale-95 duration-200 ease-in-out">Learn More</a>
                </div>
            </div>
        </section>

        <section id="business-english" class="min-h-[50vh] bg-purple-200 grid place-content-center py-32">
            <div class="hide">
                <div class="px-4 max-w-[600px] mx-auto">
                    <h2 class="font-bold text-deep text-5xl text-balance md:text-center mb-8">Professional English classes for career success.</h2>
                    <p class="text-lg text-gray-800 mb-4">Our classes are customized for your personal needs and are designed to build well-rounded English skills and instill confidence with professional communication in English.</p>
                </div>

                <ul class="px-2 max-w-[800px] mx-auto grid grid-rows-4 sm:grid-rows-3 md:grid-rows-2 sm:grid-cols-6 md:grid-cols-8 gap-4 mb-16">
                    <li class="row-span-1 md:row-span-2 stacked rounded-xl sm:col-span-6 md:col-span-8 h-[200px]">
                        <div class="w-full h-full grid sm:grid-cols-2 md:grid-cols-3 sm:bg-purple-950 rounded-xl">
                            <div class="stacked sm:col-start-2 sm:col-span-2">
                                <img
                                    height="200"
                                    widht="200"
                                    loading="lazy"
                                    fetchpriority="low"
                                    src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-bus-eng-presentation.jpg" 
                                    class="w-full object-cover h-[200px] rounded-xl sm:rounded-tr-xl sm:rounded-br-xl"/>
                                <div class="bg-purple-950/50 sm:bg-gradient-to-r sm:from-purple-950 sm:to-transparent rounded-xl sm:rounded-tr-xl sm:rounded-br-xl"></div>
                            </div>
                        </div>
                        <div class="flex flex-col justify-end p-4">
                            <h3 class="text-2xl font-semibold text-white">Communicate with <br> Clients and Colleagues</h3>
                        </div>
                    </li>
                    <li class="row-span-1 md:row-span-2 stacked bg-white rounded-xl sm:col-span-3 md:col-span-2 h-[200px]">
                        <div class="w-full h-full">
                            <img
                                height="200"
                                wids="200"
                                loading="lazy"
                                fetchpriority="low" 
                                src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-bus-eng-write-effectively.jpg" 
                                class="w-full object-cover h-[200px] rounded-xl"/>
                        </div>
                        <div class="w-full h-full bg-purple-900/50 rounded-xl"></div>
                        <div class="flex flex-col justify-end p-4">
                            <h3 class="text-xl text-white font-semibold">Write <br> Effectively</h3>
                        </div>
                    </li>
                    <li class="row-span-1 md:row-span-2 stacked bg-white rounded-xl sm:col-span-3 md:col-span-2 h-[200px]">
                        <div class="w-full h-full">
                            <img
                                height="200"
                                width="200"
                                loading="lazy"
                                fetchpriority="low"
                                src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-bus-eng-understand.jpg"
                                class="w-full object-cover h-[200px] rounded-xl"/>
                        </div>
                        <div class="w-full h-full bg-purple-900/50 rounded-xl"></div>
                        <div class="flex flex-col justify-end p-4">
                            <h3 class="text-xl text-white font-semibold">Understand <br> Native Speakers</h3>
                        </div>
                    </li>
                    <li class="row-span-1 md:row-span-2 stacked bg-white rounded-xl sm:col-span-6 md:col-span-4 h-[200px]">
                        <div class="w-full h-full">
                            <img
                                height="200"
                                width="200"
                                loading="lazy" 
                                fetchpriority="low"
                                src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-bus-eng-master-industry.jpg"
                                class="w-full object-cover h-[200px] rounded-xl"/>
                        </div>
                        <div class="w-full h-full bg-purple-900/50 rounded-xl"></div>
                        <div class="flex flex-col justify-end p-4">
                            <h3 class="text-xl font-semibold text-white">Master <br> Industry English</h3>
                        </div>
                    </li>
                </ul>

                <div class="flex justify-center gap-4">
                    <a href="#contact" class="inline-block bg-purple-800 px-8 py-2 text-white rounded-md hover:bg-purple-700 active:bg-purple-900 active:scale-95 duration-200 ease-in-out">
                        Book Now
                    </a>
                    <a href="#business-english" class="inline-block text-lg bg-purple-300 text-purple-800 rounded-md px-4 py-1 hover:bg-purple-200 active:bg-purple-400 active:text-purple-900 active:scale-95 duration-200 ease-in-out">Learn More</a>
                </div>
            </div>
        </section>

        <section id="about">
            <div class="hide max-w-[1000px] mx-auto grid sm:grid-cols-2">
                <section class="w-full h-full">
                    <img 
                        height="500"
                        width="500"
                        loading="lazy"
                        fetchpriority="low"
                        src="https://teacher-mark.s3.ap-southeast-1.amazonaws.com/small-about-teacher.png" 
                        class="w-full h-full object-cover mx-auto"/>
                </section>
                <section class="py-8 text-center grid place-content-center">
                    <h2 class="font-bold text-deep text-4xl text-balance md:text-center mb-8 text-purple-950">About me</h2>
                    <div class="flex flex-col max-w-[200px] mx-auto gap-2">
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">7+ years experience</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">TEFL Certified</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">500+ students</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">Private classes</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">Group classes</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">Kids classes</span>
                        <span class="bg-purple-100 rounded-md px-3 py-1 text-purple-950">Adult classes</span>
                    </div>
                </section>
            </div>
        </section>

        <section id="contact" class="bg-gradient-to-r from-purple-700/80 to-purple-800/80 text-white py-24 md:py-24">
            <div class="hide max-w-[500px] mx-auto bg-white text-black px-6 py-6 sm:rounded-2xl">
                <h2 class="text-2xl font-semibold text-purple-900 mb-2">Contact Me To Arrange Your First Class</h2>
                <p>Scan the Line QR code or fill out the email form below. I am available between 8 am and 8 pm everyday to respond to your inquiries. I will do my best to reply to you within 2 hours of messaging me.</p>
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
                    <h3>Email</h3>
                    <hr />
                </div>
                {{ template "form" . }}
            </div>
        </section>

        <section id="faq">
            {{/*<div class="py-16 mx-auto px-4 sm:px-0 bg-deep text-white shadow-[0px_-30px_0px_30px_#020517] sm:shadow-[0px_-50px_0px_50px_#020517]">*/}}
            <div class="py-16 mx-auto px-4 sm:px-0 bg-deep text-white">
                <div class="hide max-w-[700px] mx-auto">
                    <h2 class="text-4xl mb-8">The answers to frequently asked questions</h2>
                    <dl>
                        <dt class="hide text-xl mb-2 cursor-pointer">How can I contact you?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">You can contact me anytime between 8AM and 8PM by Line or Email through the Contact Me section of this website.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">How long is each class?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">Students can choose from 1 hour, 1.5 hour, or 2 hour classes.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">What materials are provided by the teacher?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">Materials depend on the class. For Test Preparation classes, the teacher provides materials and students are welcome to ask for help with their own textbooks. For Business English classes, the teacher provides materials and students are welcome to bring materials related to their job or other parts of life.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">Where are the classes?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">I currently offer classes online and in-person. In-person classes can be at the student's residence, a local coffee shop, or other convenient location.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">How can I book a class?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">You can sign up for a class by contacting the teacher directly. The teacher's contact information is available in the Contact Me section of this website. The teacher is available from 8AM to 8PM by Line or Email.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">Can I book classes in advance?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">Yes, students can book single classes or multiple classes in advance. For a single class, students will choose the day with the teacher and pay in advance. For multiple classes, students will choose which days they can attend class for a lesson package. Each week, students can confirm their days with their teacher. With multiple classes, students can reschedule on 24-hour notice, because we understand that things happen in life.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">Do you offer refunds or rescheduling?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">I do offer rescheduling, but it depends on what booking the students make. If students buy multiple classes in advance, we can reschedule on 24-hour notice. However, students will need to use their classes within a specified timeframe. Refunds can be available under certain circumstances on agreement with the teacher.</dd>
                        <hr class="my-4"/>

                        <dt class="hide text-xl mb-2 cursor-pointer">How can I pay for classes?</dt>
                        <dd class="hidden max-w-[60ch] text-purple-50/90">After contacting the teacher and agreeing upon the class and time, you can pay by Line Pay or bank transfer.</dd>
                        <hr class="my-4"/>
                    </dl>
                    <p class="text-purple-50/90">For questions about other class types such as children's classes and adult conversation classes, please contact the teacher directly using the Contact Me section of the website. Teacher Mark is available from 8AM to 8PM to help you.</p>
                </div>
            </div>
        </section>

        <footer class="bg-deep text-purple-50/90 h-[48px] flex items-center justify-center">
            <p>&copy; All rights reserved Teacher Mark 2025</p>
        </footer>

        <div id="alert" class="w-[300px] min-h-6 fixed top-2 left-2 bg-green-300 shadow-lg rounded-md px-4 py-2 duration-200 -translate-y-8 opacity-0"></div>
    </main>
    <script src="static/js/behaviors/accordion.js"></script>
    <script src="static/js/behaviors/observer.js"></script>
    <script src="static/js/behaviors/toast.js"></script>
  </body>
</html>   