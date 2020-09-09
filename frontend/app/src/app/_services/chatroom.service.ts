import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Chatroom } from '../_models/chatroom'
import { Message } from '../_models/message';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {

  private chatroomsUrl = 'http://localhost:8000/messages';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

  getChatrooms(): Observable<Chatroom[]> {
    return this.http.get<Chatroom[]>(this.chatroomsUrl);
  }

  getMessages(id): Observable<Message[]> {
    return this.http.get<Message[]>(this.chatroomsUrl+'/'+id);
  }

  sendMessage(id, message): Observable<Message[]> {
    return this.http.post<Message[]>(this.chatroomsUrl+'/'+id, {message});
  }
}
