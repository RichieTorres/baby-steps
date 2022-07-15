import 'package:flutter/material.dart';
import 'package:navigation/screens/one.dart';
import 'package:navigation/screens/three.dart';
import 'package:navigation/screens/two.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Navigation',
      routes: {
        '/one': (context) => ScreenOne(),
        '/two': (context) => ScreenTwo(),
        '/three': (context) => ScreenThree()
      },
      initialRoute: '/one',
    );
  }
}
