import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ProfileComponent} from './profile/profile.component';
import {LoginComponent} from './login/login.component';
import { AuthGuard } from './_helpers/auth.guard';
import { GetMatchedComponent} from './get-matched/get-matched.component';
import { FeedComponent } from './feed/feed.component';
import { ChatsComponent } from './chats/chats.component';
import { EmptyChatsComponent } from './empty-chats/empty-chats.component';

const routes: Routes = [
  { path: '', component: ProfileComponent, canActivate: [AuthGuard]},
  { path: 'profile', component: ProfileComponent, canActivate: [AuthGuard]},
  { path: 'profile/:id', component: ProfileComponent},
  { path: 'login', component: LoginComponent },
  { path: 'chats', component: EmptyChatsComponent },
  { path: 'chats/:id', component: ChatsComponent },
  { path: 'feed', component: FeedComponent },
  { path: 'get-matched', component: GetMatchedComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
