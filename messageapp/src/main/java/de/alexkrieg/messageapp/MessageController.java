package de.alexkrieg.messageapp;

import java.util.Date;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

import de.alexkrieg.messageapp.domain.Message;

@Controller
class MessageController {

    @Autowired
    MessageSender messageSender;

    @GetMapping("/form")
    public String formGet(Model model) {
        model.addAttribute("message", new Message());
        return "message";
    }

    @PostMapping("/form")
    public String formPost(Message message, BindingResult bindingResult, Model model) {
        message.setCreationTime(new Date());
        message.setAuthor("messageapp");

        messageSender.sendMessage(message);
        if (!bindingResult.hasErrors()) {
            model.addAttribute("noErrors", true);
        }
        model.addAttribute("message", message);
        return "message";
    }

    @GetMapping(path = "/logout")
    public String logout(HttpServletRequest request) throws ServletException {
        request.logout();
        return "http://bigdemo.com";
    }
}