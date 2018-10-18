<#import "/spring.ftl" as spring />

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Message creation</title>
        <link href="/css/main.css" rel="stylesheet">
    </head>
    <body>
        <h2>Your message</h2>

        <@spring.bind "message"/>
        
        <#if message?? && noErrors??>
            Your submitted message:
            <div>${message.text}</div><br>
            <a href="form">New Message</a>
        <#else>
            <form action="form" method="post">
                Text:<br>
                <@spring.formInput "message.text"/>
                <@spring.showErrors "<br>"/>
                <input type="submit" value="Submit">
            </form>
        </#if>

        <a href="${baseUrl}">Back to UI</a>
        <script src="/js/main.js"></script>
    </body>
</html>