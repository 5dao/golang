var FindProxyForURL = function(init, profiles) {
    return function(url, host) {
        "use strict";
        var result = init, scheme = url.substr(0, url.indexOf(":"));
        do {
            result = profiles[result];
            if (typeof result === "function") result = result(url, host, scheme);
        } while (typeof result !== "string" || result.charCodeAt(0) === 43);
        return result;
    };
}("+auto switch", {
    "+auto switch": function(url, host, scheme) {
        "use strict";
        if (/(?:^|\.)qnap\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)amazonaws\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)amazon\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)cloudfront\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)google\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)golang\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)mktoresp\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)zoom\.us$/.test(host)) return "+v2ray";
        if (/(?:^|\.)yes-news\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)zaobao\.com\.sg$/.test(host)) return "+v2ray";
        if (/(?:^|\.)zaobao\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)sphdigital\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)scorecardresearch\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)outbrain\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)nextmedia\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)pao-pao\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)secretchina\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)dongtaiwang\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)cnd\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)huobi\.co$/.test(host)) return "+v2ray";
        if (/(?:^|\.)quantcount\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)doubleclick\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)facebook\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)crazyegg\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)googletagservices\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)quantserve\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)hk01\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)dwnews\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)dwnews\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)githubassets\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bitcointalk\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)6parknews\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)boxun\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)aboluowang\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)creaders\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)chinadigitaltimes\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)telesco\.pe$/.test(host)) return "+v2ray";
        if (/(?:^|\.)telegram\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)t\.me$/.test(host)) return "+v2ray";
        if (/(?:^|\.)jquery\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)nokia\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)googleadservices\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)google-analytics\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ctfassets\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)go-acme\.github\.io$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ajax\.googleapis\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)content-people-pa\.googleapis\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)content\.googleapis\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)gstatic\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)mitbbs\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)recaptcha\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)contexts\.co$/.test(host)) return "+v2ray";
        if (/(?:^|\.)twitter\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)brawersoftware\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)facebook\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)split\.io$/.test(host)) return "+v2ray";
        if (/(?:^|\.)godaddy\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)wsimg\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)gravatar\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)pubmine\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)tumblr\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)wp\.com$/.test(host)) return "+v2ray";
        if (/^gist\.github\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)openwrt\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)wenxuecity\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)trafficmanager\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)youtube\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)epochtimes\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)backchina\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)rfi\.fr$/.test(host)) return "+v2ray";
        if (/(?:^|\.)binance\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)zb\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)duping\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bitmex\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)poloniex\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)mxpnl\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)twoeggz\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)googleusercontent\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)akamaized\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)airmailapp\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bwh1\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)fbcdn\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)android\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bluekai\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)semasio\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)servebom\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)disqus\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)powerlinks\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)howtogeek\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)googlesyndication\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)vmware\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)wootric\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)d2wy8f7a9ursnm\.cloudfront\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)d36jcksde1wxzq\.cloudfront\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)github-production-release-asset-2e65be\.s3\.amazonaws\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)carbonads\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)d2dfho4r6t7asi\.cloudfront\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)speakerdeck\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)taboola\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)nytcn\.me$/.test(host)) return "+v2ray";
        if (/(?:^|\.)line-scdn\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)nyt\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)nytimes\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)reddit\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)wikipedia\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)medium\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)rfa\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)quora\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)zynbit\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ads-twitter\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)iterm2colorschemes\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)unpkg\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)whatsapp\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)twimg\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)adnxs\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)dw\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bbc\.co\.uk$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bbci\.co\.uk$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bbc\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ggpht\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ytimg\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)googlevideo\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ssr\.tools$/.test(host)) return "+v2ray";
        if (/(?:^|\.)v2ray\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)toutyrater\.github\.io$/.test(host)) return "+v2ray";
        if (/(?:^|\.)google\.com\.hk$/.test(host)) return "+v2ray";
        if (/(?:^|\.)imgur\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)sitescout\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)bttrack\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)adsrvr\.org$/.test(host)) return "+v2ray";
        if (/(?:^|\.)ml314\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)sourceforge\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)microsoft\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)clicktale\.net$/.test(host)) return "+v2ray";
        if (/(?:^|\.)office365\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)whatsapp\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)trustarc\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)oraclecloud\.com$/.test(host)) return "+v2ray";
        if (/(?:^|\.)amazonaws\.com$/.test(host)) return "+v2ray";
        return "DIRECT";
    },
    "+v2ray": function(url, host, scheme) {
        "use strict";
        if (/^127\.0\.0\.1$/.test(host) || /^::1$/.test(host) || /^localhost$/.test(host)) return "DIRECT";
        return "SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080";
    }
});