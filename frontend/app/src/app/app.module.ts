import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

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

import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { FeedImageComponent } from './feed/feed-image/feed-image.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatListModule } from '@angular/material/list';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';


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
    FeedImageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    MatIconModule,
    MatCardModule,
    MatGridListModule,
    MatListModule,
    NoopAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
