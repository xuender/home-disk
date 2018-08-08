import { NgModule, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http'
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { ComponentsModule } from '../components/components.module'
import { PipesModule } from '../pipes/pipes.module'
import { IonicImageViewerModule } from 'ionic-img-viewer'
import { MyApp } from './app.component';

import { AboutPage } from '../pages/about/about';
import { ContactPage } from '../pages/contact/contact';
import { FilesPage } from '../pages/files/files';
import { TabsPage } from '../pages/tabs/tabs';
import { ImagePage } from '../pages/image/image'

import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { FilesProvider } from '../providers/files/files';
import { PreviewProvider } from '../providers/preview/preview';
import { InfoPage } from '../pages/info/info';
import { UploadPage } from '../pages/upload/upload';

@NgModule({
  declarations: [
    MyApp,
    AboutPage,
    ContactPage,
    FilesPage,
    ImagePage,
    TabsPage,
    InfoPage,
    UploadPage,
  ],
  imports: [
    BrowserModule,
    ComponentsModule,
    PipesModule,
    HttpClientModule,
    IonicImageViewerModule,
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    AboutPage,
    ContactPage,
    FilesPage,
    TabsPage,
    ImagePage,
    UploadPage,
    InfoPage
  ],
  providers: [
    StatusBar,
    SplashScreen,
    { provide: ErrorHandler, useClass: IonicErrorHandler },
    FilesProvider,
    PreviewProvider
  ]
})
export class AppModule { }
