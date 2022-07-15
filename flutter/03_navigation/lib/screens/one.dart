import 'package:flutter/material.dart';
import 'package:navigation/shared/menu_drawer.dart';

import '../shared/menu_bottom.dart';

class ScreenOne extends StatelessWidget {
  const ScreenOne({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Screen One'),
      ),
      drawer: const MenuDrawer(),
      bottomNavigationBar: const MenuBottom(),
      body: const Center(child: Text('1')),
    );
  }
}
