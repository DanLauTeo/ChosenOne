import { Injectable } from '@angular/core';
import { Observable, of, empty, throwError } from 'rxjs';
import { catchError, map, tap, flatMap } from 'rxjs/operators';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Chatroom } from '../_models/chatroom'
import { Message } from '../_models/message';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {

  basePath: string = "/api/v1"
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private httpClient: HttpClient) { }

  startChat(userId: string): Observable<Chatroom> {
    return this.httpClient.post<Chatroom>(`${this.basePath}/chatrooms/`, {"requested_user_id": userId})
      .pipe(
        catchError((error, _) => {
          if (error instanceof HttpErrorResponse && error.status == 409) {
            return this.getChatrooms()
              .pipe(
                map((chatrooms, _) => {
                  return chatrooms.find((chatroom, _) => chatroom.participants.includes(userId));
                })
              );
          } else {
            return throwError(error);
          }
        })
      )
  }

  getChatrooms(): Observable<Chatroom[]> {
    return this.httpClient.get<Chatroom[]>(`${this.basePath}/chatrooms/`);
  }

  getChatroom(id: number) {
    return this.httpClient.get<Chatroom>(`${this.basePath}/chatrooms/${id}/`)
  }

  getMessages(id: number): Observable<Message[]> {
    return this.httpClient.get<Message[]>(`${this.basePath}/chatrooms/${id}/messages/`);
  }

  sendMessage(id: number, message: string): Observable<Message> {
    return this.httpClient.post<Message>(`${this.basePath}/chatrooms/${id}/messages/`, {"message": message});
  }

  getMessage(id: number): Observable<Message> {
    return this.httpClient.get<Message>(`${this.basePath}/messages/${id}/`)
  }

  deleteMessage(id: number) {
    return this.httpClient.delete(`${this.basePath}/messages/${id}/`)
  }
}
