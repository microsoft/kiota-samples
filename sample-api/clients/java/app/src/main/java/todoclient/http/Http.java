// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package todoclient.http;

import java.io.InputStream;
import java.security.KeyStore;
import java.security.SecureRandom;
import java.util.Properties;

import javax.net.ssl.SSLContext;
import javax.net.ssl.TrustManagerFactory;
import javax.net.ssl.X509TrustManager;

import com.microsoft.kiota.http.KiotaClientFactory;

import okhttp3.OkHttpClient;

public class Http {
    private static final String PROPERTIES_FILE = "devcert.properties";
    private static final String CERT_FILE_KEY = "certfile";
    private static final String CERT_PASSWORD_KEY = "certpassword";

    // This creates a client that will accept the dev cert used
    // by the sample API project. See README for more information
    public static OkHttpClient getHttpClientWithDevCert() {
        try {
            final Properties certProperties = new Properties();
            certProperties.load(Http.class.getResourceAsStream(PROPERTIES_FILE));

            final String certFile = certProperties.getProperty(CERT_FILE_KEY);
            final String certPassword = certProperties.getProperty(CERT_PASSWORD_KEY);
            final KeyStore keyStore = KeyStore.getInstance("PKCS12");

            final InputStream pfxFile = Http.class.getResourceAsStream(certFile);

            keyStore.load(pfxFile, certPassword.toCharArray());
            pfxFile.close();

            final TrustManagerFactory trustManagerFactory = TrustManagerFactory.getInstance("X509");
            trustManagerFactory.init(keyStore);

            final SSLContext sslContext = SSLContext.getInstance("SSL");
            sslContext.init(null, trustManagerFactory.getTrustManagers(), new SecureRandom());

            return KiotaClientFactory.create()
                .sslSocketFactory(sslContext.getSocketFactory(), (X509TrustManager)trustManagerFactory.getTrustManagers()[0])
                .build();
        } catch (Exception e) {
            return null;
        }
    }
}
