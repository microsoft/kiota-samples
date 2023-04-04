import { config } from '../config/config';
import { Logger } from './common/logger';
require('node-fetch');

export class App {
  static async run(): Promise<App> {
    const app = new App();
    app.start();
    Logger.logTask('hellow world');
    return app;
  }

  private start(): void {
    this.logAppInfo();
  }

  private logAppInfo(): void {
    Logger.logTask('APP', {
      develop: DEVELOP,
      version: VERSION,
      config: config,
    });
  }
}
