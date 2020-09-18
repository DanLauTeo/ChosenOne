import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError, map, tap, flatMap } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Chatroom } from '../_models/chatroom'
import { Message } from '../_models/message';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private httpClient: HttpClient) { }

  startChat(userId: string): Observable<Chatroom> {
    return this.httpClient.post<Chatroom>(`/chatrooms/`, {"requested_user_id": userId})
  }

  getChatrooms(): Observable<Chatroom[]> {
    return this.httpClient.get<Chatroom[]>(`/chatrooms/`);
  }

  getChatroom(id: number) {
    return this.httpClient.get<Chatroom>(`/chatrooms/${id}/`)
  }

  getMessages(id: number): Observable<Message[]> {
    return this.httpClient.get<Message[]>(`/chatrooms/${id}/messages/`);
  }

  sendMessage(id: number, message: string): Observable<Message> {
    return this.httpClient.post<Message>(`/chatrooms/${id}/messages/`, {"message": message});
  }

  getMessage(id: number): Observable<Message> {
    return this.httpClient.get<Message>(`/messages/${id}/`)
  }

  deleteMessage(id: number) {
    return this.httpClient.delete(`/messages/${id}/`)
  }
}
