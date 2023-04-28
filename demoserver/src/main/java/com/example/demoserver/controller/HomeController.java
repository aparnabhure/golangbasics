package com.example.demoserver.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.*;

import java.util.Map;


@Controller
@Slf4j
@RequestMapping(value = "/v1/home")
@RequiredArgsConstructor
public class HomeController {


    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseStatus(value = HttpStatus.OK)
    @ResponseBody
    public ResponseEntity<Object> home() {
        return ResponseEntity.ok("Welcome");
    }

    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseStatus(value = HttpStatus.OK)
    @ResponseBody
    public ResponseEntity<Object> postHome(@RequestHeader Map<String, String> headers, @RequestBody Map<String, Object> request) {
        return ResponseEntity.ok(request);
    }

    @PostMapping(value = "/form", consumes = MediaType.APPLICATION_FORM_URLENCODED_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseStatus(value = HttpStatus.OK)
    @ResponseBody
    public ResponseEntity<Object> postForm(@RequestHeader Map<String, String> headers, @RequestParam MultiValueMap<String,String> paramMap) {
        return ResponseEntity.ok(paramMap);
    }

}

