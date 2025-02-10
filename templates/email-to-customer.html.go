<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Info</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: black;
            color: white;
            padding: 20px;
        }
        .container {
            background-color: #fff;
            border-radius: 8px;
            padding: 20px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        }

        .text-container {
            background-color: black;
            color: white;
            padding: 1rem;
        }
        .footer {
            font-size: 12px;
            text-align: center;
            margin-top: 20px;
        }
    </style>
</head>
<body style="background-color: #8D50BD; padding: 4rem 0;">
     <div style="background-color: #fff; color: rgb(19, 18, 20); padding: 1rem;">
        <div style="max-width: 600px; margin: 0 auto;">
            <h1 style="color: #8D50BD; font-size: 24px;">
            {{ if eq .Lang "en" }}
                Contact Form Submission
            {{ else }}
                聯絡表單提交
            {{ end }}
            </h1>
            <p>
            {{ if eq .Lang "en" }}
                {{.Name}},
            {{ else }}
                姓名 {{.Name}},
            {{ end }}
            </p>

            <p>
            {{ if eq .Lang "en" }}
            Thank you for reaching out about English classes! I'm excited to learn how I can help you achieve your English goals.
            {{ else }}
            謝謝您對英語課程的興趣！我很高興了解如何幫助您實現英語目標。
            {{ end }}
            </p>
            <p>
            {{ if eq .Lang "en" }}
            I will get back to you within 24 hours to schedule a free consultation. In the meantime, feel free to keep reading our <a href="https://teachermark.com.tw/en">website</a> for more information about our classes.
            {{ else }}
            我將在24小時內與您聯繫，安排免費諮詢。與此同時，請隨時瀏覽我們的<a href="https://teachermark.com.tw">網站</a>，以獲取有關我們課程的更多信息。
            {{ end }}
            </p>
            <p>
            {{ if eq .Lang "en" }}
                Kindest regards,
            {{ else }}
                致以最誠摯的問候，
            {{ end }}
            </p>
            <p>
            {{ if eq .Lang "en" }}
                Teacher Mark
            {{ else }}
                老師馬克
            {{ end }}
            </p>
            <hr style="border: 1px solid #8D50BD; margin: 4rem 0 2rem;">
            <footer class="footer">
                <p>&copy; | Teacher Mark 2025 <br> All Rights Reserved</p>
            </footer>
        </div>
    </div>
</body>
</html>