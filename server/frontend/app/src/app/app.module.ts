import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';
import { TopBarComponent } from './top-bar/top-bar.component';
import { GalleryComponent } from './gallery/gallery.component';
import { FeedComponent } from './feed/feed.component';
import { ChatsComponent } from './chats/chats.component';
import { GetMatchedComponent } from './get-matched/get-matched.component';
import { MatchComponent } from './get-matched/match/match.component';
import { ProfilePicComponent } from './profile/profile-pic/profile-pic.component';
import { ChatlistComponent } from './chats/chatlist/chatlist.component';
import { MessagesComponent } from './chats/messages/messages.component';
import { DeleteMessageDialogComponent } from './chats/messages/delete-message-dialog/delete-message-dialog.component';
import { EmptyChatsComponent } from './empty-chats/empty-chats.component';
import { PopupComponent } from './gallery/popup/popup.component';
import { FeedImageComponent } from './feed/feed-image/feed-image.component';

import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    ProfileComponent,
    TopBarComponent,
    GalleryComponent,
    FeedComponent,
    ChatsComponent,
    GetMatchedComponent,
    MatchComponent,
    ProfilePicComponent,
    ChatlistComponent,
    MessagesComponent,
    EmptyChatsComponent,
    FeedImageComponent,
    PopupComponent,
    DeleteMessageDialogComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NoopAnimationsModule,
    MatDividerModule,
    MatListModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    MatGridListModule,
    MatDialogModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
