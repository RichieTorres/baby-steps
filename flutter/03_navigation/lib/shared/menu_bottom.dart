import 'package:flutter/material.dart';

class MenuBottom extends StatelessWidget {
  const MenuBottom({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: const [
        BottomNavigationBarItem(
            icon: Icon(Icons.looks_one_rounded), label: 'One'),
        BottomNavigationBarItem(
            icon: Icon(Icons.looks_two_rounded), label: 'Two'),
        BottomNavigationBarItem(
            icon: Icon(Icons.looks_3_rounded), label: 'Three')
      ],
      onTap: (int index) {
        switch (index) {
          case 0:
            Navigator.pushNamed(context, '/one');
            break;
          case 1:
            Navigator.pushNamed(context, '/two');
            break;
          case 2:
            Navigator.pushNamed(context, '/three');
            break;
        }
      },
    );
  }
}
