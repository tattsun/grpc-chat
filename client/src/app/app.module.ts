import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { JoinRequest, JoinResponse, SayRequest, SayResponse } from '../api/chat_pb.d'
import { Chat } from '../api/chat_pb_service'

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
      BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
