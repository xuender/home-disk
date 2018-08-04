import { NgModule, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { ComponentsModule } from '../components/components.module'
import { PipesModule } from '../pipes/pipes.module'
import { MyApp } from './app.component';

import { AboutPage } from '../pages/about/about';
import { ContactPage } from '../pages/contact/contact';
import { FilesPage } from '../pages/files/files';
import { HomePage } from '../pages/home/home';
import { TabsPage } from '../pages/tabs/tabs';

import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { FilesProvider } from '../providers/files/files';

@NgModule({
  declarations: [
    MyApp,
    AboutPage,
    ContactPage,
    FilesPage,
    HomePage,
    TabsPage
  ],
  imports: [
    BrowserModule,
    ComponentsModule,
    PipesModule,
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    AboutPage,
    ContactPage,
    FilesPage,
    HomePage,
    TabsPage
  ],
  providers: [
    StatusBar,
    SplashScreen,
    { provide: ErrorHandler, useClass: IonicErrorHandler },
    FilesProvider
  ]
})
export class AppModule { }
