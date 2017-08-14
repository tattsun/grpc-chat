import { Component } from '@angular/core';
import { grpc, Code, Metadata } from 'grpc-web-client';
import { JoinRequest, JoinResponse, SayRequest, SayResponse } from '../api/chat_pb'
import { Chat } from '../api/chat_pb_service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
    title = 'app';

    constructor() {
        console.log("unko");

        const joinRequest = new JoinRequest();
        joinRequest.setName("John");

        grpc.unary(Chat.Join, {
            request: joinRequest,
            host: "https://localhost:5000",
            onEnd: res => {
                console.log(res);
            }
        });
    }
}
