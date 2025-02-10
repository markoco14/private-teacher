{{ define "form" }}
<form 
    hx-post='{{ if eq .Lang "en"}} /en/contact {{ else }} /contact {{ end }}'
    hx-swap="outerHTML"
    hx-disabled-elt="find button"
    action='{{ if eq .Lang "en"}} /en/contact {{ else }} /contact {{ end }}'
    method="post">
    <div class="flex flex-col mb-2">
        <label for="name" class="text-purple-900">
        {{ if eq .Lang "en" }}
            Name
        {{ else }}
            姓名
        {{ end }}
        </label>
        <p class="text-gray-700 text-sm mb-1">
        {{ if eq .Lang "en" }}
            Share your name so we know how to address you.
        {{ else }}
            請告訴我們您的名字，以便我們知道如何稱呼您。
        {{ end }}
        </p>
        <input 
            id="name"
            type="text"
            name="name"
            class="border rounded p-2 w-full"
            placeholder='{{ if eq .Lang "en" }} Ms. Emily {{ else }} 陳小姐 {{ end }}'
            value="{{.Form.Name}}"/>
        <p class="h-5 text-right text-sm">
            {{range .Errors}}
                {{if eq .Field "name"}}
                    <span class="text-red-500">{{.Message}}</span>
                {{end}}
            {{end}}
        </p>
    </div>
    <div class="flex flex-col mb-2">
        <label for="email" class="text-purple-900">
        {{ if eq .Lang "en" }}
            Email
        {{ else }}
            電子郵件
        {{ end }}
        </label>
        <p class="text-gray-700 text-sm mb-1">
        {{ if eq .Lang "en" }}
            Use the email where you want us to contact you.
        {{ else }}
            使用您希望我們與您聯繫的電子郵件。
        {{ end }}
        </p>
        <input 
            id="email"
            type="text"
            name="email"
            class="border rounded p-2 w-full"
            placeholder="your-email@email.com"
            value="{{.Form.Email}}"/>
        <p class="h-5 text-right text-sm">
            {{range .Errors}}
                {{if eq .Field "email"}}
                    <span class="text-red-500">{{.Message}}</span>
                {{end}}
            {{end}}
        </p>
    </div>
    <div class="flex flex-col mb-2">
        <label for="message" class="text-purple-900">
        {{ if eq .Lang "en" }}
            Message
        {{ else }}
            訊息
        {{ end }}
        </label>
        <p class="text-gray-700 text-sm mb-1">
        {{ if eq .Lang "en" }}
            Write a short message to say what type of class you are looking for.
        {{ else }}
            寫一個簡短的消息，說明您想要的課程類型。
        {{ end }}
        </p>
        <textarea 
            id="message"
            name="message"
            class="border rounded p-2 w-full" 
            rows="5">{{.Form.Message}}</textarea>
        <p class="h-5 text-right text-sm">
            {{range .Errors}}
                {{if eq .Field "message"}}
                    <span class="text-red-500">{{.Message}}</span>
                {{end}}
            {{end}}
        </p>
    </div>
    <button 
        type="submit"
        hx-indicator="this"
        class="bg-purple-700 text-white rounded px-4 py-1 self-start w-full hover:bg-purple-600 active:bg-purple-800 active:scale-95 cursor-pointer duration-200 ease-in-out disabled:bg-purple-900 disabled:opacity-80 disabled:cursor-default disabled:active:bg-purple-900 disabled:active:scale-100" >
        <span class="htmx-indicator-hide">
        {{ if eq .Lang "en" }}
            Send message
        {{ else }}
            發送消息
        {{ end }}
        </span>
        <span class="htmx-indicator">
        {{ if eq .Lang "en" }}
            Sending...
        {{ else }}
            發送中...
        {{ end }}
        </span>
    </button>
</form>
{{ end }}