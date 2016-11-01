package com.openatx.apkparser;

import java.io.File;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import com.google.gson.Gson;

import Decoder.BASE64Encoder;
import net.dongliu.apk.parser.ApkParser;
import net.dongliu.apk.parser.bean.ApkMeta;
import net.dongliu.apk.parser.bean.Icon;

public class Main {

	public static void main(String[] args) {

		Map<String, Object> configMap = new HashMap<String, Object>();

		if (args.length != 1 ) {
			System.out.println("Args number should be 1.");
			System.exit(0);
		}

		String apkPath = args[0];

		if (apkPath.contains("\\.") && apkPath.split("\\.")[1] != "apk") {
			System.out.println("Bad package name.");
			System.exit(0);
		}

		ApkParser apkParser;
		try {
			apkParser = new ApkParser(new File(apkPath));
			ApkMeta apkMeta = apkParser.getApkMeta();

			// Get Config
			configMap.put("packageName", apkMeta.getPackageName());
			configMap.put("label", apkMeta.getLabel());
			configMap.put("versionName", apkMeta.getVersionName());
			configMap.put("versionCode", apkMeta.getVersionCode());
			configMap.put("minSdkVersion", apkMeta.getMinSdkVersion());
			configMap.put("targetSdkVersion", apkMeta.getTargetSdkVersion());
			configMap.put("maxSdkVersion", apkMeta.getMaxSdkVersion());

			Icon icon = apkParser.getIconFile();
			configMap.put("iconPath", icon.getPath());

			BASE64Encoder encoder = new BASE64Encoder();
			configMap.put("icon", encoder.encode(icon.getData()));

			Gson gson = new Gson();
			String json = gson.toJson(configMap);

			System.out.println(json);

		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}
